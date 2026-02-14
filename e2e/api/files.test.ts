import { expect, test } from "../src/fixtures";

test.describe("Files", () => {
  test.describe("List files", () => {
    test("should return 200", async ({ request }) => {
      const response = await request.get("/api/v1/files");

      expect(response.status()).toBe(200);
    });

    test("should return 200 for version id", async ({ createProject, createVersion, request }) => {
      const project = await createProject();
      const version = await createVersion({ projectId: project.id });

      const response = await request.get("/api/v1/files", {
        params: {
          versionId: version.id,
        }
      });

      expect(response.status()).toBe(200);
    });

    test("should return 400 for invalid version id", async ({ request }) => {
      const response = await request.get("/api/v1/files", {
        params: {
          versionId: 'invalid-id',
        }
      });

      expect(response.status()).toBe(400);
    });
  });

  test.describe("Create files", () => {
    test("should return 201", async ({ request }) => {
      const response = await request.post("/api/v1/files", {
        data: {
          name: "example.txt"
        },
      });

      expect(response.status()).toBe(201);
    });

    test("should return 400 for missing name", async ({ request }) => {
      const response = await request.post("/api/v1/files", {
        data: {},
      });

      expect(response.status()).toBe(400);
    });
  });

  test.describe("Upload file", () => {
    test("should return 201", async ({ createFile, request }) => {
      const file = await createFile({ name: "example.txt" });

      const response = await request.post(`/api/v1/files/${file.id}/upload`, {
        multipart: {
          file: {
            name: "example.txt",
            mimeType: "text/plain",
            buffer: Buffer.from("Hello, world!")
          }
        }
      });

      expect(response.status()).toBe(201);

      const getFileResponse = await request.get(`/api/v1/files/${file.id}`);

      expect(getFileResponse.status()).toBe(200);
      await expect(getFileResponse.json()).resolves.toEqual(expect.objectContaining({
        name: "example.txt",
        mimeType: "text/plain; charset=utf-8",
        size: 13,
        isComplete: true
      }));
    });

    test("should return 400 for missing file", async ({ createFile, request }) => {
      const file = await createFile({ name: "example.txt" });

      const response = await request.post(`/api/v1/files/${file.id}/upload`, {
        data: {},
      });

      expect(response.status()).toBe(400);
    });

    test("should return 409 for already uploaded file", async ({ createFile, request }) => {
      const file = await createFile({ name: "example.txt" });

      const firstResponse = await request.post(`/api/v1/files/${file.id}/upload`, {
        multipart: {
          file: {
            name: "example.txt",
            mimeType: "text/plain",
            buffer: Buffer.from("Hello, world!")
          }
        }
      });

      expect(firstResponse.status()).toBe(201);

      const secondResponse = await request.post(`/api/v1/files/${file.id}/upload`, {
        multipart: {
          file: {
            name: "example.txt",
            mimeType: "text/plain",
            buffer: Buffer.from("Hello, world!")
          }
        }
      });

      expect(secondResponse.status()).toBe(409);
    });

    test("should return 400 for invalid file ID", async ({ request }) => {
      const response = await request.post(`/api/v1/files/invalid-id/upload`);

      expect(response.status()).toBe(400);
    });

    test("should return 400 for non-existing file", async ({ request }) => {
      const response = await request.post(`/api/v1/files/-1/upload`);

      expect(response.status()).toBe(400);
    });
  });

  test.describe("Download file", () => {
    test("should return 200", async ({ createFile, request }) => {
      const file = await createFile({
        name: "example.txt",
        mimeType: "text/plain",
        buffer: Buffer.from("Hello, world!")
      });

      const response = await request.get(`/api/v1/files/${file.id}/download`);

      expect(response.status()).toBe(200);
      expect(response.headers()['content-type']).toBe('text/plain; charset=utf-8');
      await expect(response.text()).resolves.toBe('Hello, world!');
    });

    test("should return 404 for incomplete file", async ({ createFile, request }) => {
      const file = await createFile({ name: "example.txt" });

      const response = await request.get(`/api/v1/files/${file.id}/download`);

      expect(response.status()).toBe(404);
    });

    test("should return 400 for invalid file ID", async ({ request }) => {
      const response = await request.get(`/api/v1/files/invalid-id/download`);

      expect(response.status()).toBe(400);
    });

    test("should return 404 for non-existing file", async ({ request }) => {
      const response = await request.get(`/api/v1/files/-1/download`);

      expect(response.status()).toBe(404);
    });
  });

  test.describe("Get file", () => {
    test("should return 200", async ({ createFile, request }) => {
      const file = await createFile({ name: "example.txt" });

      const response = await request.get(`/api/v1/files/${file.id}`);

      expect(response.status()).toBe(200);
    });

    test("should return 400 for invalid file ID", async ({ request }) => {
      const response = await request.get(`/api/v1/files/invalid-id`);

      expect(response.status()).toBe(400);
    });

    test("should return 404 for non-existing file", async ({ request }) => {
      const response = await request.get(`/api/v1/files/-1`);

      expect(response.status()).toBe(404);
    });
  });

  test.describe("Delete file", () => {
    test("should return 204", async ({ createFile, request }) => {
      const file = await createFile({ name: "example.txt" });

      const response = await request.delete(`/api/v1/files/${file.id}`);

      expect(response.status()).toBe(204);
    });

    test("should return 400 for invalid version ID", async ({ request }) => {
      const response = await request.delete(`/api/v1/files/invalid-id`);

      expect(response.status()).toBe(400);
    });

    test("should return 404 for non-existing version", async ({ request }) => {
      const response = await request.delete(`/api/v1/files/-1`);

      expect(response.status()).toBe(404);
    });
  });
});
