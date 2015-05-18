class InvitationsController < ApplicationController
  # GET /invitations/new
  def new
    respond_to do |format|
      format.html{ render 'invitations/_new', layout: false }
    end
  end

  # POST /invitations
  def create
  end

  # GET /invitations/:id/accept
  def accept
  end

  # GET /invitations/:id/decline
  def decline
  end
end
