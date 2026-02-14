import { test as baseTest, mergeExpects } from "@playwright/test";
import {
  createCreateProjectFixture,
  CreateProjectParams,
  CreateProjectResult,
} from "./fixtures/project";
import { createCreateVersionFixture, CreateVersionParams, CreateVersionResult } from "./fixtures/version";
import { createCreateFileFixture, CreateFileParams, CreateFileResult } from "./fixtures/file";

type TestFixtures = {
  createProject: (params?: CreateProjectParams) => Promise<CreateProjectResult>;
  createVersion: (params: CreateVersionParams) => Promise<CreateVersionResult>;
  createFile: (params?: CreateFileParams) => Promise<CreateFileResult>;
};

export const test = baseTest.extend<TestFixtures>({
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
});

export const expect = mergeExpects();
