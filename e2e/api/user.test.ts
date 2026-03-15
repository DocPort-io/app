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
  });

  test.describe("User", () => {
    test.describe("Create", () => {
      test("should return 200", async ({ request, defaultToken }) => {
        const response = await request.post("/api/v1/users", {
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

      test("should return 400 for missing name", async ({ request, defaultToken }) => {
        const response = await request.post("/api/v1/users", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            email: `john-${uuid.v4()}@example.com`,
            emailVerified: true
          }
        });

        expect(response.status()).toBe(400);
      });

      test("should return 400 for missing email", async ({ request, defaultToken }) => {
        const response = await request.post("/api/v1/users", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            name: `John - ${uuid.v4()}`,
            emailVerified: true
          }
        });

        expect(response.status()).toBe(400);
      });

      test("should return 409 for duplicate email", async ({ request, defaultToken }) => {
        const email = `john-${uuid.v4()}@example.com`;

        await request.post("/api/v1/users", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            name: `John - ${uuid.v4()}`,
            email,
            emailVerified: true
          }
        });

        const response = await request.post("/api/v1/users", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          },
          data: {
            name: `John - ${uuid.v4()}`,
            email,
            emailVerified: true
          }
        });

        expect(response.status()).toBe(409);
      });
    });

    test.describe("Information", () => {
      test("should return 200 for valid user", async ({ request, defaultToken }) => {
        const createResponse = await request.post("/api/v1/users", {
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
  });
});
