class UsersController < ApplicationController
	before_action :authenticate_user!
	load_and_authorize_resource

	# GET /users/:id
	#
	# Displays the profile information of a particular user
	#
	# @param [Integer] id Identifier of the user to be displayed
	def show
	end
end
