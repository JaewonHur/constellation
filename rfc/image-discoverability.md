# OS image & measurement discovery

The Constellation OS image build pipeline generates a set of images using a chosen commit of the Constellation monorepo and and a desired release version number.

```mermaid
graph LR
    version["input: version (<code>v2.2.0</code>)"] --> imageid["image version uid (<code>v2.2.0</code>)"]
    commit["input: Constellation repo commit hash (<code>cc0de5c</code>)"] --> pipeline
    imageid --> pipeline{build images}
    pipeline --> awsimage["raw AWS image"]
    pipeline --> azureimage[raw Azure image]
    pipeline --> gcpimage[raw GCP image]
    awsimage --> awsupload{AWS ami upload}
    azureimage --> azureupload{Azure image upload}
    gcpimage --> gcpupload{GCP image upload}
    gcpimage --> gcpmeasure{measure}
    awsimage --> awsmeasure{measure}
    azureimage --> azuremeasure{measure}
    awsupload --> awsimage1[ami-123]
    awsupload --> awsimage2[ami-456]
    awsupload --> awsimage3[ami-789]
    azureupload --> azureimage1[azure-cvm-123]
    azureupload --> azureimage2[azure-trusted-launch-123]
    gcpupload --> gcpimage1[gcp-image-123]
    awsmeasure --> awsmeasurements[measurements.yml]
    azuremeasure --> azuremeasurements[measurements.yml]
    gcpmeasure --> gcpmeasurements[measurements.yml]
    awsimage1 --> lut{Upload image lookup table to S3<br>key: <code>v2.2.0</code><br>value: lookup table for image references specific to each CSP}
    awsimage2 --> lut
    awsimage3 --> lut
    azureimage1 --> lut
    azureimage2 --> lut
    gcpimage1 --> lut
    awsmeasurements --> s3measurements{"Upload image measurements to S3<br><code>aws/v2.2.0/measurements.yaml<br>azure/v2.2.0/measurements.yaml<br>gcp/v2.2.0/measurements.yaml<br></code>"}
    azuremeasurements --> s3measurements
    gcpmeasurements --> s3measurements
```

## Inputs and outputs of the build pipeline

The build pipeline takes as inputs:

- a version number that is one of
  - a release version number (e.g. `v2.2.0`) for release images
  - a pseudo-version number (e.g. `v2.3.0-pre-debug`) for development images
  - a pseudo-version number (e.g. `v2.3.0-pre-branch-name`) for branch images
- a commit hash of the Constellation monorepo that is used to build the images (e.g. `cc0de5c68d41f31dd0b284d574f137e0b0ad106b`)

To identify images belonging to one invocation of the build pipeline, the pipeline uses a unique identifier for the set of images, referred to as `image version uid`.
This is either the release version number (e.g. `v2.2.0`) or a pseudo version that combines the version number and the commit hash (e.g. `v2.3.0-pre-debug-cc0de5c68d41f31dd0b284d574f137e0b0ad106b`).

The build pipeline produces as outputs:

- a raw OS image for every CSP
- a set of measurements for each raw OS image
- one or more images uploaded to each CSP (e.g. AWS AMIs, Azure images, GCP images)
- a lookup table that maps the `image version uid` to the CSP-specific image references

The lookup table is uploaded to S3 and is used to identify the images that belong to a given `image version uid`.
Measurements are uploaded to S3 and can be looked up for each cloud service provider and `image version uid`.

## Image lookup table

The image lookup table is a JSON file that maps the `image version uid` to the CSP-specific image references. It uses the `image version uid` as file name.

```
s3://<BUCKET-NAME>/images/<IMAGE-VERSION-UID>.json
```

```json
{
  "aws": {
    "us-east-1": "ami-123",
    "us-west-2": "ami-456",
    "eu-west-1": "ami-789"
  },
  "azure": {
    "cvm": "azure-cvm-123",
    "trustedlaunch": "azure-trusted-launch-123"
  },
  "gcp": {
    "sev-es": "gcp-image-123"
  }
  "qemu": "https://cdn.edgeless.systems/.../qemu.raw"
}
```

- For AWS, the image lookup table contains the AMI IDs for each region
- For Azure, the image lookup table contains the image IDs for the CVM and Trusted Launch images
- For GCP, the image lookup table contains the image ID
- For QEMU, the image lookup table contains a URL to the QEMU image

This document is not signed and can be extended in the future to include more image references (e.g. if an image is replicated to a new AWS region).
The same document can be used to identify old images that are no longer used and can be deleted for cost optimization.

## Image measurements

This RFC is not about the image measurements themselves, but about how they are stored and looked up.
The format of the image measurements is described in the [secure software distribution RFC](secure-software-distribution.md).

The image measurements are stored in a folder structure in S3 that is organized by CSP and `image version uid`.

```
s3://<BUCKET-NAME>/measurements/<CSP>/<IMAGE-VERSION-UID>/measurements.yaml
s3://<BUCKET-NAME>/measurements/<CSP>/<IMAGE-VERSION-UID>/measurements.yaml.sig
```

## CLI image discovery

The CLI needs to be able to discover the image references for a given `image version uid`.
By default, the CLI will prefill the `image` field of the `constellation-conf.yaml` when `constellation config generate <CSP>` is run with a hardcoded `image version uid` (e.g. `v2.2.0`).
The `image` field is independent of the CSP and is a used to discover the CSP-specific image reference as needed for the following operations:

- `constellation create`
- `constellation upgrade apply`

The CLI can find a CSP- and region specific image reference by looking up the `image version uid` in the following order:

- if a local file `<IMAGE-VERSION-UID>.json` exists, use the lookup table in that file
- otherwise, load the image lookup table from a well known URL (e.g. `https://cdn.confidential.cloud/images/<IMAGE-VERSION-UID>.json`) and use the lookup table in that file
- choose the CSP-specific image reference for the current region and security type:
  - On AWS, use the AMI ID for the current region (e.g. `.aws.us-east-1`)
  - On Azure, use the image ID for the security type (CVM or Trusted Launch) (e.g. `.azure.cvm`)
  - On GCP, use the only image ID (e.g. `.gcp`)

This allows customers to upload images to their own cloud subscription and use them with the CLI by providing the image lookup table as a local file.

## Future extensions

This is a list of possible future extensions that are not part of this RFC.
Their implementation is not guaranteed.
They are listed here to ensure that the design of this RFC is flexible enough to support them.

- A lookup table for available image versions might be added in the future.
- The lookup table can be signed using a signing key that is only used for that purpose.
- User managed repositories can be added in the future. This would allow users to reupload Constellation OS images to their cloud subscription and host their own lookup tables that resolve the same image versions to image references pointing to self managed images. An optional `repository` field could be added to the configuration file to allow users to specify the repository to use for image discovery.