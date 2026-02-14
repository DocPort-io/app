import { APIRequestContext } from "@playwright/test";
import { expect } from "../fixtures";

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
    return {
      accessToken: responseBody.access_token,
    };
  };

  return {
    authenticate,
  };
};
