import { APIRequestContext } from "@playwright/test";
import { expect } from "../fixtures";
import { jwtDecode } from "jwt-decode";

export type AuthenticateResult = {
  accessToken: string;
};

export type AuthenticateParams = {
  clientId?: string;
  clientSecret?: string;
  username?: string;
  password?: string;
};

export const createAuthenticateFixture = (request: APIRequestContext) => {
  const authenticate = async (params: AuthenticateParams): Promise<AuthenticateResult> => {
    const response = await request.post(process.env.KEYCLOAK_TOKEN_URL, {
      form: {
        grant_type: 'password',
        client_id: params.clientId,
        client_secret: params.clientSecret,
        username: params.username,
        password: params.password,
      }
    });

    expect(response.status()).toBe(200);

    const responseBody = await response.json();

    const decodedToken = jwtDecode(responseBody.access_token);

    // Wait for the token to be valid
    while (decodedToken.iat > (Date.now() / 1000)) {
      await new Promise((resolve) => setTimeout(resolve, 100));
    }

    return {
      accessToken: responseBody.access_token,
    };
  };

  return {
    authenticate,
  };
};
