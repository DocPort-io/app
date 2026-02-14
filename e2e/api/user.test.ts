import { expect, test } from "../src/fixtures";
import * as uuid from "uuid";

test.describe("Users", () => {
  test.describe("Me", () => {
    test.describe("Information", () => {
      test("should return 404 for unknown user", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/me", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(404);
      });
    });

    test.describe("External Auths", () => {
      test("should return 404 for unknown user", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/me/external-auths", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(404);
      });
    });
  });

  test.describe("User", () => {
    test.describe("Create", () => {
      test("should return 200", async ({ request, defaultToken }) => {
        const response = await request.post("/api/v1/users/", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            name: `John - ${uuid.v4()}`,
            email: `john-${uuid.v4()}@example.com`,
            emailVerified: true
          }
        });

        expect(response.status()).toBe(200);
      });
    });

    test.describe("Information", () => {
      test("should return 200 for valid user", async ({ request, defaultToken }) => {
        const createResponse = await request.post("/api/v1/users/", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            name: `John - ${uuid.v4()}`,
            email: `john-${uuid.v4()}@example.com`,
            emailVerified: true
          }
        });

        const createBody = await createResponse.json();

        const response = await request.get(`/api/v1/users/${createBody.id}`, {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(200);
      });

      test("should return 404 for unknown user id", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/-1", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(404);
      });

      test("should return 404 for invalid user id", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/invalid-id", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(400);
      });
    });

    test.describe("External Auths", () => {
      test("should return 200 for valid user", async ({ request, defaultToken }) => {
        const createResponse = await request.post("/api/v1/users/", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            name: `John - ${uuid.v4()}`,
            email: `john-${uuid.v4()}@example.com`,
            emailVerified: true
          }
        });

        const createBody = await createResponse.json();

        const response = await request.get(`/api/v1/users/${createBody.id}/external-auths`, {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(200);
      });

      test.fail("should return 404 for unknown user id", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/-1/external-auths", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(404);
      });

      test("should return 404 for invalid user id", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/invalid-id/external-auths", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(400);
      });
    });

    test.describe("External Auths - Create", () => {
      test("should return 201", async ({ request, defaultToken }) => {
        const createResponse = await request.post("/api/v1/users/", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            name: `John - ${uuid.v4()}`,
            email: `john-${uuid.v4()}@example.com`,
            emailVerified: true
          }
        });

        const createBody = await createResponse.json();

        const response = await request.post(`/api/v1/users/${createBody.id}/external-auths`, {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            provider: "example",
            providerId: uuid.v4()
          }
        });

        expect(response.status()).toBe(201);
      });
    });
  });
});
