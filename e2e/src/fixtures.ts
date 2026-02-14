import { mergeExpects, test as baseTest } from "@playwright/test";
import { createCreateProjectFixture, CreateProjectParams, CreateProjectResult, } from "./fixtures/project";
import { createCreateVersionFixture, CreateVersionParams, CreateVersionResult } from "./fixtures/version";
import { createCreateFileFixture, CreateFileParams, CreateFileResult } from "./fixtures/file";
import { AuthenticateParams, AuthenticateResult, createAuthenticateFixture } from "./fixtures/auth";

type TestFixtures = {
  tokenUrl: string;
  defaultClientId: string;
  defaultUsername: string;
  defaultPassword: string;
  defaultToken: string;
  createProject: (params?: CreateProjectParams) => Promise<CreateProjectResult>;
  createVersion: (params: CreateVersionParams) => Promise<CreateVersionResult>;
  createFile: (params?: CreateFileParams) => Promise<CreateFileResult>;
  authenticate: (params: AuthenticateParams) => Promise<AuthenticateResult>;
};

export const test = baseTest.extend<TestFixtures>({
  tokenUrl: async ({}, use) => {
    await use(process.env.KEYCLOAK_TOKEN_URL);
  },
  defaultClientId: async ({}, use) => {
    await use(process.env.KEYCLOAK_CLIENT_ID);
  },
  defaultUsername: async ({}, use) => {
    await use(process.env.KEYCLOAK_USERNAME);
  },
  defaultPassword: async ({}, use) => {
    await use(process.env.KEYCLOAK_PASSWORD);
  },
  defaultToken: async ({ authenticate, defaultClientId, defaultUsername, defaultPassword }, use) => {
    const result = await authenticate({
      clientId: defaultClientId,
      username: defaultUsername,
      password: defaultPassword,
    });
    await use(result.accessToken);
  },
  createProject: async ({ request }, use) => {
    const { createProject } = createCreateProjectFixture(request);
    await use(createProject);
  },
  createVersion: async ({ request }, use) => {
    const { createVersion } = createCreateVersionFixture(request);
    await use(createVersion);
  },
  createFile: async ({ request }, use) => {
    const { createFile } = createCreateFileFixture(request);
    await use(createFile);
  },
  authenticate: async ({ request }, use) => {
    const { authenticate } = createAuthenticateFixture(request);
    await use(authenticate);
  }
});

export const expect = mergeExpects();
