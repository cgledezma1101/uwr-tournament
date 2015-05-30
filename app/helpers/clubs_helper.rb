module ClubsHelper
  # Allows the creation of a method that will render management actions over a club's admins
  #
  # @param [Club] club The club for which the actions will be performed
  #
  # @return [Proc] A method that can be called with a particular administrator and that will render the actions that can be performed on that administrator for the mentioned club
  def create_admin_actions_proc(club)
    Proc.new do |user|
      club_admin = ClubAdmin.find_by(club: club, user: user)
      render 'clubs/admin_actions', club_admin: club_admin
    end
  end

  # Allows the creation of a method that will render management actions over club members
  #
  # @param [Club] club The club for which the actions will be received
  #
  # @return [Proc] A method that can be called with a particular member and that will render the actions that can be performed on that member for the mentioned club
  def create_member_actions_proc(club)
    Proc.new do |user|
      user_club = UserClub.find_by(club: club, user: user)
      render 'clubs/member_actions', user_club: user_club
    end
  end
end
