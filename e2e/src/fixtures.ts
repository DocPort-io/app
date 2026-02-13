import { test as baseTest, mergeExpects } from "@playwright/test";
import {
  createCreateProjectFixture,
  CreateProjectParams,
  CreateProjectResult,
} from "./fixtures/project";
import { createCreateVersionFixture, CreateVersionParams, CreateVersionResult } from "./fixtures/version";

type TestFixtures = {
  createProject: (params?: CreateProjectParams) => Promise<CreateProjectResult>;
  createVersion: (params: CreateVersionParams) => Promise<CreateVersionResult>;
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
});

export const expect = mergeExpects();
