class Users::OmniauthCallbacksController < Devise::OmniauthCallbacksController
  def facebook
    auth =
    {
      name: request.env['omniauth.auth'].info.name,
      email: request.env['omniauth.auth'].info.email,
      uid: request.env['omniauth.auth'].uid,
      provider: request.env['omniauth.auth'].provider
    }

    user = User.from_omniauth(auth)
    if(user.persisted?)
      sign_in_and_redirect user
    else
      redirect_to new_user_session_path, alert: t('devise.omniauth.sign_in_error')
    end
  end
end
