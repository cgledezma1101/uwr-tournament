class Ability
  include CanCan::Ability

  def initialize(user)
    can :show, Club do |club|
      user.clubs.where{ club_id == my{club.id} }.any?
    end
  end
end
