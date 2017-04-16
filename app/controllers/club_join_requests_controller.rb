class ClubJoinRequestsController < ApplicationController
	before_action :authenticate_user!
	load_and_authorize_resource

	# GET /club_join_requests/:id/accept
	#
	# @param [Integer] id Identifier of the membership request to accept
	def accept
		join_request = ClubJoinRequest.find(params[:id])
		membership = UserClub.new(club: join_request.club, user: join_request.user)

		if(membership.save)
			join_request.destroy
			redirect_params = { notice: t('club_join_request.accept_success', name: membership.user.name) }
		else
			redirect_params = { alert: t('club_join_request.accept_failure') }
		end

		redirect_to club_path(membership.club), redirect_params
	end

	# POST /club_join_requests
	def create
		@club_join_request.user = current_user
		if(@club_join_request.save)
			redirect_params = { notice: t('club_join_request.create_success') }
		else
			redirect_params = { alert: t('club_join_request.create_fail') }
		end
		redirect_to user_path(current_user), redirect_params
	end

	# GET /club_join_requests/:id/decline
	#
	#@param [Integer] id Identifier of the membership request to decline
	def decline
		join_request = ClubJoinRequest.find(params[:id])
		join_request.destroy
		redirect_to club_path(join_request.club), notice: t('club_join_request.declined', user_name: join_request.user.name)
	end

	# GET /club_join_requests/new
	def new
		@available_clubs = Club.all.to_a - current_user.all_clubs - current_user.pending_clubs.to_a - current_user.pending_join_requests.to_a
		@available_clubs.uniq!
		@available_clubs.sort! do |club0, club1|
			club0.name <=> club1.name
		end
		render 'club_join_requests/_new', layout: false
	end

	private

	def create_params
		params.require(:club_join_request).permit(:club_id)
	end
end
