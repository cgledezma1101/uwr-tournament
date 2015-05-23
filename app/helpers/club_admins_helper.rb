module ClubAdminsHelper
  # Allows the creation of a method that will render specia actions for administrators
  #
  # @param [Club] club The club for whom the actions will be received
  #
  # @return [Proc] A method that can be called with a particular administrator and that will render the actions that can be performed on that administrator for the mentioned club
  def create_admin_actions_proc(club)
    Proc.new do |user|
      club_admin = ClubAdmin.find_by(club: club, user: user)
      render 'clubs/admin_actions', club_admin: club_admin
    end
  end
end
