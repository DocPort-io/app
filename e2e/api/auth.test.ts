import { expect, test } from "../src/fixtures";

test.describe("Auth", () => {
  test("should return 401 for invalid token", async ({ request }) => {
    const response = await request.get("/api/v1/users/me/token-info", {
      headers: {
        "Authorization": "Bearer token"
      }
    });

    expect(response.status()).toBe(401);
  });

  test("should return 401 for missing token", async ({ request }) => {
    const response = await request.get("/api/v1/users/me/token-info");

    expect(response.status()).toBe(401);
  });

  test("should return 401 for invalid token prefix", async ({ request }) => {
    const response = await request.get("/api/v1/users/me/token-info", {
      headers: {
        "Authorization": "Unknown token"
      }
    });

    expect(response.status()).toBe(401);
  });

  test("should return 200 for valid token", async ({
                                                     authenticate,
                                                     request,
                                                     defaultClientId,
                                                     defaultUsername,
                                                     defaultPassword
                                                   }) => {
    const auth = await authenticate({
      clientId: defaultClientId,
      username: defaultUsername,
      password: defaultPassword,
    });

    const response = await request.get("/api/v1/users/me/token-info", {
      headers: {
        "Authorization": `Bearer ${auth.accessToken}`
      }
    });

    expect(response.status()).toBe(200);
    await expect(response.json()).resolves.toEqual(expect.objectContaining({
      sub: expect.any(String)
    }));
  });
});
