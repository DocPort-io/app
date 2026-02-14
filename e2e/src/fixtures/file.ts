import { APIRequestContext } from "@playwright/test";
import { expect } from "../fixtures";

export type CreateFileResult = {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  size: string | null;
  mimeType: string | null;
  isComplete: boolean;
};

export type CreateFileParams = {
  name?: string;
  mimeType?: string;
  buffer?: Buffer;
};

export const createCreateFileFixture = (request: APIRequestContext) => {
  const createFile = async (params: CreateFileParams): Promise<CreateFileResult> => {
    const createResponse = await request.post("/api/v1/files", {
      data: {
        name: params.name ?? "example.txt"
      },
    });

    expect(createResponse.status()).toBe(201);

    const createResponseBody = await createResponse.json();

    if (params.mimeType && params.buffer) {
      const uploadResponse = await request.post(`/api/v1/files/${createResponseBody.id}/upload`, {
        multipart: {
          file: {
            name: params.name,
            mimeType: params.mimeType,
            buffer: params.buffer,
          }
        }
      });

      expect(uploadResponse.status()).toBe(201);
    }

    const response = await request.get(`/api/v1/files/${createResponseBody.id}`);

    expect(response.status()).toBe(200);

    const responseBody = await response.json();
    return responseBody;
  };

  return {
    createFile,
  };
};
