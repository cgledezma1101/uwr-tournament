class InvitationsController < ApplicationController
  before_action :authenticate_user!

  # GET /invitations/new?club_id=:club_id&is_admin=:is_admin
  #
  # @param [Integer] club_id Identifier of the club the invitation is being created for
  def new
    club = Club.find(params[:club_id])
    authorize! :edit, club

    @invitation = Invitation.new(club: club, is_admin: params[:is_admin])
    render 'invitations/_new', layout: false
  end

  # POST /invitations
  def create
    club = Club.find(params[:invitation][:club_id])
    authorize! :edit, club

    user = User.find_by(email: params[:user][:email])

    if(user.nil?)
      invitation = Invitation.create(club: club)
    else
      invitation = Invitation.create(club: club, user: user)
    end

    if(invitation.valid?)
      redirect_to club_path(club), notice: t('invitation.create_success')
    else
      redirect_to club_path(club), alert: t('invitation.create_failure')
    end
  end

  # POST /invitations/:id/accept
  #
  # Allows a user, who has been previously invited, to join a club
  #
  # @param [Integer] id Identifier of the invitation that's being accepted
  def accept
    invitation = Invitation.find(params[:id])
    authorize! :accept, invitation

    membership = UserClub.new(club: invitation.club, user: invitation.user)
    if(membership.save)
      invitation.destroy
      redirect_to club_path(membership.club), notice: I18n.t('invitation.accepted', club_name: membership.club_name)
    else
      redirect_to user_path(invitation.user), alert: I18n.t('invitation.accept_error')
    end
  end

  # POST /invitations/:id/decline
  #
  # Allows a user to reject an invitation that has been produced for her, by a particular club
  #
  # @param [Integer] id Identifier of the invitation being rejected
  def decline
    invitation = Invitation.find(params[:id])
    authorize! :decline, invitation

    invitation.destroy
    redirect_to user_path(invitation.user), notice: t('invitation.declined', club_name: invitation.club_name)
  end
end
