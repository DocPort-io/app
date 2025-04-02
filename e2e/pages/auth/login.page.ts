import { BasePage } from '../../framework/base.page';

export class LoginPage extends BasePage {
	readonly href = '/auth/login';

	readonly inputEmail = this.page.getByTestId('login-email');
	readonly inputPassword = this.page.getByTestId('login-password');
	readonly buttonSignIn = this.page.getByTestId('login-signin');
	readonly textErrorTitle = this.page.getByTestId('login-error-title');
	readonly textErrorDescription = this.page.getByTestId('login-error-description');
}
