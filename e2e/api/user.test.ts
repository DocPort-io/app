import { expect, test } from "../src/fixtures";

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
    test.describe("Information", () => {
      test("should return 404 for unknown user", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/-1", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(404);
      });
    });

    test.describe("External Auths", () => {
      test.fail("should return 404 for unknown user", async ({ request, defaultToken }) => {
        const response = await request.get("/api/v1/users/-1/external-auths", {
          headers: {
            "Authorization": `Bearer ${defaultToken}`
          }
        });

        expect(response.status()).toBe(404);
      });
    });
  })
});
