class ApplicationController < ActionController::Base
	# Prevent CSRF attacks by raising an exception.
	# For APIs, you may want to use :null_session instead.
	protect_from_forgery with: :exception
	before_action :configure_permitted_parameters, if: :devise_controller?

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
		devise_parameter_sanitizer.for(:sign_up) << :name
	end
end
