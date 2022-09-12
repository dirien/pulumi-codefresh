import * as pulumi from "@pulumi/pulumi";
import * as codefresh from "@pulumiverse/codefresh"

const test = new codefresh.Account("test", {
  builds: [{
    parallel: 27,
  }],
  features: {
    OfflineLogging: true,
    abac: true,
    customKubernetesCluster: true,
    launchDarklyManagement: false,
    ssoManagement: true,
    teamsManagement: true,
  },
  limits: [{
    collaborators: 25,
    dataRetentionWeeks: 5,
  }],
});

export const testName = test.name
