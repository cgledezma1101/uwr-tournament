class Users::OmniauthCallbacksController < Devise::OmniauthCallbacksController
  def facebook
    sign_in_user(request.env['omniauth.auth'])
  end

  def google_oauth2
    sign_in_user(request.env['omniauth.auth'])
  end

  private

  # Standard way of signing in from an OmniAuth response
  #
  # @param Hash auth Hash returned by OmniAuth when the provider responded successfully
  def sign_in_user(auth)
    user = User.from_omniauth(auth)
    if(user.persisted?)
      sign_in_and_redirect user
    else
      redirect_to new_user_session_path, alert: t('devise.omniauth.sign_in_error')
    end
  end
end
