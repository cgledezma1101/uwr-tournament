class InvitationsController < ApplicationController
  # GET /invitations/new?club_id=:club_id&is_admin=:is_admin
  #
  # @param [Integer] club_id Identifier of the club the invitation is being created for
  # @param [Boolean] is_admin Boolean indicating whether this is an invitation to be administrator
  def new
    club = Club.find(params[:club_id])
    authorize! :edit, club

    @invitation = Invitation.new(club: club, is_admin: params[:is_admin])

    respond_to do |format|
      format.html{ render 'invitations/_new', layout: false }
    end
  end

  # POST /invitations
  def create
    club = Club.find(params[:invitation][:club_id])
    authorize! :edit, club

    user = User.find_by(email: params[:user][:email])

    if(user.nil?)
      invitation = Invitation.create(club: club, is_admin: params[:invitation][:is_admin])
    else
      invitation = Invitation.create(club: club, user: user, is_admin: params[:invitation][:is_admin])
    end

    if(invitation.valid?)
      redirect_to club_path(club), notice: t('invitations.create_success')
    else
      redirect_to club_path(club), alert: t('invitations.create_failure')
    end
  end

  # GET /invitations/:id/accept
  def accept
  end

  # GET /invitations/:id/decline
  def decline
  end
end