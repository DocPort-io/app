import { APIRequestContext } from "@playwright/test";
import * as uuid from "uuid";
import { expect } from "../fixtures";

export type CreateProjectResult = {
  id: number;
  createdAt: string;
  updatedAt: string;
  slug: string;
  name: string;
};

export type CreateProjectParams = {
  slug?: string;
  name?: string;
};

export const createCreateProjectFixture = (request: APIRequestContext) => {
  const createProject = async (params: CreateProjectParams = {}) => {
    const response = await request.post("/api/v1/projects", {
      data: {
        slug: params.slug ?? uuid.v4(),
        name: params.name ?? "Test project",
      },
    });

    expect(response.status()).toBe(201);

    const responseBody = await response.json();
    return responseBody;
  };

  return {
    createProject,
  };
};
