import { test, expect } from "../src/fixtures";

test.describe("Heartbeat", () => {
  test("should return 200", async ({ request }) => {
    const response = await request.get("/heartbeat");
    expect(response.status()).toBe(200);
  });
});
