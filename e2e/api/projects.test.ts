import { expect, test } from "../src/fixtures";
import * as uuid from "uuid";

test.describe("Projects", () => {
  test.describe("List projects", () => {
    test("should return 200", async ({ request }) => {
      const response = await request.get("/api/v1/projects");

      expect(response.status()).toBe(200);
    });
  });

  test.describe("Create project", () => {
    test("should return 201", async ({ request }) => {
      const response = await request.post("/api/v1/projects", {
        data: {
          slug: uuid.v4(),
          name: "Test project",
        },
      });

      expect(response.status()).toBe(201);
    });

    test("should return 400 for missing name", async ({ request }) => {
      const response = await request.post("/api/v1/projects", {
        data: {
          slug: uuid.v4(),
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 400 for missing slug", async ({ request }) => {
      const response = await request.post("/api/v1/projects", {
        data: {
          name: "Test project",
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 409 for duplicate slug", async ({ createProject, request }) => {
      const project = await createProject();

      const response = await request.post("/api/v1/projects", {
        data: {
          slug: project.slug,
          name: "Test project",
        },
      });

      expect(response.status()).toBe(409);
    });
  });

  test.describe("Get project", () => {
    test("should return 200", async ({ createProject, request }) => {
      const project = await createProject();

      const response = await request.get(`/api/v1/projects/${project.id}`);

      expect(response.status()).toBe(200);
    });

    test("should return 400 for invalid project ID", async ({ request }) => {
      const response = await request.get(`/api/v1/projects/invalid-id`);

      expect(response.status()).toBe(400);
    });

    test("should return 404 for non-existing project", async ({ request }) => {
      const response = await request.get(`/api/v1/projects/-1`);

      expect(response.status()).toBe(404);
    });
  });

  test.describe("Update project", () => {
    test("should return 200", async ({ createProject, request }) => {
      const project = await createProject();

      const response = await request.put(`/api/v1/projects/${project.id}`, {
        data: {
          slug: uuid.v4(),
          name: "Updated project name",
        },
      });

      expect(response.status()).toBe(200);
    });

    test("should return 400 for invalid project ID", async ({ request }) => {
      const response = await request.put(`/api/v1/projects/invalid-id`, {
        data: {
          slug: uuid.v4(),
          name: "Updated project name",
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 400 for missing name", async ({ createProject, request }) => {
      const project = await createProject();

      const response = await request.put(`/api/v1/projects/${project.id}`, {
        data: {
          slug: uuid.v4(),
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 400 for missing slug", async ({ createProject, request }) => {
      const project = await createProject();

      const response = await request.put(`/api/v1/projects/${project.id}`, {
        data: {
          name: "Updated project name",
        },
      });

      expect(response.status()).toBe(400);
    });

    test("should return 404 for non-existing project", async ({ request }) => {
      const response = await request.put(`/api/v1/projects/-1`, {
        data: {
          slug: uuid.v4(),
          name: "Updated project name",
        },
      });

      expect(response.status()).toBe(404);
    });
  });

  test.describe("Delete project", () => {
    test("should return 204", async ({ createProject, request }) => {
      const project = await createProject();

      const response = await request.delete(`/api/v1/projects/${project.id}`);

      expect(response.status()).toBe(204);
    });

    test("should return 400 for invalid project ID", async ({ request }) => {
      const response = await request.delete(`/api/v1/projects/invalid-id`);

      expect(response.status()).toBe(400);
    });

    test.fail("should return 404 for non-existing project", async ({ request }) => {
      const response = await request.delete(`/api/v1/projects/-1`);

      expect(response.status()).toBe(404);
    });
  });
});
