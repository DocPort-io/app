import { expect, test } from "../src/fixtures";
import { CreateProjectResult } from "../src/fixtures/project";

test.describe("Versions", () => {
  let project: CreateProjectResult;

  test.beforeEach(async ({ createProject }) => {
    project = await createProject();
  });

  test.describe("List versions by project ID", () => {
    test("should return 200", async ({ request }) => {
      const response = await request.get("/api/v1/versions", {
        params: {
          projectId: project.id,
        },
      });

      expect(response.status()).toBe(200);
    });
  });

  test.describe("Create version", () => {
    test("should return 201", async ({ request }) => {
      const response = await request.post("/api/v1/versions", {
        data: {
          name: "Test version",
          description: "Test description",
          projectId: project.id,
        },
      });

      expect(response.status()).toBe(201);
    });

    test("should return 400 for missing name", async ({ request }) => {
      const response = await request.post("/api/v1/versions", {
        data: {
          description: "Test description",
          projectId: project.id,
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 201 with a missing description", async ({ request }) => {
      const response = await request.post("/api/v1/versions", {
        data: {
          name: "Test version",
          projectId: project.id,
        },
      });

      expect(response.status()).toBe(201);
    });

    test("should return 400 for missing project ID", async ({ request }) => {
      const response = await request.post("/api/v1/versions", {
        data: {
          name: "Test version",
          description: "Test description",
        },
      });

      expect(response.status()).toBe(400);
    });
  });

  test.describe("Get project", () => {
    test("should return 200", async ({ createVersion, request }) => {
      const version = await createVersion({ projectId: project.id });

      const response = await request.get(`/api/v1/versions/${version.id}`);

      expect(response.status()).toBe(200);
    });

    test("should return 400 for invalid version ID", async ({ request }) => {
      const response = await request.get(`/api/v1/versions/invalid-id`);

      expect(response.status()).toBe(400);
    });

    test("should return 404 for non-existing version", async ({ request }) => {
      const response = await request.get(`/api/v1/versions/-1`);

      expect(response.status()).toBe(404);
    });
  });

  test.describe("Update version", () => {
    test("should return 200", async ({ createVersion, request }) => {
      const version = await createVersion({ projectId: project.id });

      const response = await request.put(`/api/v1/versions/${version.id}`, {
        data: {
          name: "Updated version name",
          description: "Updated version description",
        },
      });

      expect(response.status()).toBe(200);
    });

    test("should return 400 for invalid version ID", async ({ request }) => {
      const response = await request.put(`/api/v1/versions/invalid-id`, {
        data: {
          name: "Updated version name",
          description: "Updated version description",
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 400 for missing name", async ({ createVersion, request }) => {
      const version = await createVersion({ projectId: project.id });

      const response = await request.put(`/api/v1/versions/${version.id}`, {
        data: {
          description: "Updated version description",
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 200 with missing description", async ({ createVersion, request }) => {
      const version = await createVersion({ projectId: project.id });

      const response = await request.put(`/api/v1/versions/${version.id}`, {
        data: {
          name: "Updated version name",
        },
      });

      expect(response.status()).toBe(200);
    });

    test("should return 404 for non-existing version", async ({ request }) => {
      const response = await request.put(`/api/v1/versions/-1`, {
        data: {
          name: "Updated version name",
          description: "Updated version description",
        },
      });

      expect(response.status()).toBe(404);
    });
  });

  test.describe("Delete version", () => {
    test("should return 204", async ({ createVersion, request }) => {
      const version = await createVersion({ projectId: project.id });

      const response = await request.delete(`/api/v1/versions/${version.id}`);

      expect(response.status()).toBe(204);
    });

    test("should return 400 for invalid version ID", async ({ request }) => {
      const response = await request.delete(`/api/v1/versions/invalid-id`);

      expect(response.status()).toBe(400);
    });

    test.fail("should return 404 for non-existing version", async ({ request }) => {
      const response = await request.delete(`/api/v1/versions/-1`);

      expect(response.status()).toBe(404);
    });
  });
});
