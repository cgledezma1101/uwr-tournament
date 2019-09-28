class ApplicationController < ActionController::Base
    # Prevent CSRF attacks by raising an exception.
    # For APIs, you may want to use :null_session instead.
    protect_from_forgery with: :exception
    before_action :configure_permitted_parameters, if: :devise_controller?

    around_action :with_timezone

    rescue_from CanCan::AccessDenied do
        redirect_to root_url, alert: t('cancan.unauthorized_access')
    end

    def stringify_errors(model_errors, model_name)
        grouped_errors = model_errors.keys.map do |error_property|
            joined_errors = model_errors[error_property].join(',')
            translated_property = I18n.t(model_name + '.' + error_property.to_s)
            return "#{translated_property}: #{joined_errors}"
        end

        return grouped_errors.join('. ')
    end

    private

    def configure_permitted_parameters
        devise_parameter_sanitizer.permit(:sign_up, keys: [:name])
    end

    def with_timezone
        user_timezone = Time.find_zone(cookies["UWR-User-Timezone"])
        Time.use_zone(user_timezone) { yield }
    end
end
