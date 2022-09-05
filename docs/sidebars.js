/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  // By default, Docusaurus generates a sidebar from the docs folder structure
  // tutorialSidebar: [{type: 'autogenerated', dirName: '.'}],

  // But you can create a sidebar manually
  docs: [
    {
      type: 'doc',
      label: 'Welcome to Constellation',
      id: 'intro'
    },
    {
      type: 'category',
      label: 'Overview',
      link: {
        type: 'generated-index',
      },
      items: [
        {
          type: 'doc',
          label: 'Confidential Kubernetes',
          id: 'overview/confidential-kubernetes',
        },
        {
          type: 'doc',
          label: 'Security benefits',
          id: 'overview/security-benefits',
        },
        {
          type: 'doc',
          label: 'Product features',
          id: 'overview/product',
        },
        {
          type: 'doc',
          label: 'Feature status of clouds',
          id: 'overview/clouds',
        },
        {
          type: 'doc',
          label: 'Performance',
          id: 'overview/performance',
        },
        {
          type: 'doc',
          label: 'License',
          id: 'overview/license',
        },
      ]
    },
    {
      type: 'category',
      label: 'Getting started',
      link: {
        type: 'generated-index',
      },
      items: [
        {
          type: 'doc',
          label: 'Install the CLI',
          id: 'getting-started/install',
        },
        {
          type: 'doc',
          label: 'First steps',
          id: 'getting-started/first-steps',
        },
        {
          type: 'category',
          label: 'Examples',
          link: {
            type: 'doc',
            id: 'getting-started/examples',
          },
          items: [
            {
              type: 'doc',
              label: 'Emojivoto',
              id: 'getting-started/examples/emojivoto'
            },
            {
              type: 'doc',
              label: 'Online Boutique',
              id: 'getting-started/examples/online-boutique'
            },
            {
              type: 'doc',
              label: 'Horizontal Pod Autoscaling',
              id: 'getting-started/examples/horizontal-scaling'
            },
          ]
        },
      ],
    },
    {
      type: 'category',
      label: 'Workflows',
      link: {
        type: 'generated-index',
      },
      items: [
        {
          type: 'doc',
          label: 'Create your cluster',
          id: 'workflows/create',
        },
        {
          type: 'doc',
          label: 'Verify your cluster',
          id: 'workflows/verify',
        },
        {
          type: 'doc',
          label: 'Scale your cluster',
          id: 'workflows/scale',
        },
        {
          type: 'doc',
          label: 'Upgrade your cluster',
          id: 'workflows/upgrade',
        },
        {
          type: 'doc',
          label: 'Use persistent storage',
          id: 'workflows/storage',
        },
        {
          type: 'doc',
          label: 'Managing SSH keys',
          id: 'workflows/ssh',
        },
        {
          type: 'doc',
          label: 'Troubleshooting',
          id: 'workflows/troubleshooting',
        },
        {
          type: 'doc',
          label: 'Terminate your cluster',
          id: 'workflows/terminate',
        },
        {
          type: 'doc',
          label: 'Recover your cluster',
          id: 'workflows/recovery',
        },
      ],
    },
    {
      type: 'category',
      label: 'Architecture',
      link: {
        type: 'generated-index',
      },
      items: [
        {
          type: 'doc',
          label: 'Overview',
          id: 'architecture/overview',
        },
        {
          type: 'doc',
          label: 'Components',
          id: 'architecture/components',
        },
        {
          type: 'doc',
          label: 'Attestation',
          id: 'architecture/attestation',
        },
        {
          type: 'doc',
          label: 'Cluster orchestration',
          id: 'architecture/orchestration',
        },
        {
          type: 'doc',
          label: 'Versions and support',
          id: 'architecture/versions',
        },
        {
          type: 'doc',
          label: 'Images',
          id: 'architecture/images',
        },
        {
          type: 'doc',
          label: 'Keys and cryptographic primitives',
          id: 'architecture/keys',
        },
        {
          type: 'doc',
          label: 'Encrypted persistent storage',
          id: 'architecture/encrypted-storage',
        },
        {
          type: 'doc',
          label: 'Networking',
          id: 'architecture/networking',
        },
      ],
    },
    {
      type: 'category',
      label: 'Reference',
      link: {
        type: 'generated-index',
      },
      items: [
        {
          type: 'doc',
          label: 'CLI',
          id: 'reference/cli',
        },
        {
          type: 'doc',
          label: 'Configuration file',
          id: 'reference/config',
        },
      ],
    },
  ],
};

module.exports = sidebars;