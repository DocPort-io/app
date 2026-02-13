import { APIRequestContext } from "@playwright/test";
import * as uuid from "uuid";
import { expect } from "../fixtures";

export type CreateVersionResult = {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  description: string;
  projectId: number;
};

export type CreateVersionParams = {
  name?: string;
  description?: string;
  projectId: number;
};

export const createCreateVersionFixture = (request: APIRequestContext) => {
  const createVersion = async (params: CreateVersionParams): Promise<CreateVersionResult> => {
    const response = await request.post("/api/v1/versions", {
      data: {
        name: params.name ?? "Test version",
        description: params.description ?? "Test description",
        projectId: params.projectId,
      },
    });

    expect(response.status()).toBe(201);

    const responseBody = await response.json();
    return responseBody;
  };

  return {
    createVersion,
  };
};
