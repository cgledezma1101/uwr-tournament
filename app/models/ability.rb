class Ability
  include CanCan::Ability

  def initialize(user)

    ###################################################
    ################## CLUBS ##########################
    ###################################################
    can :show, Club do |club|
      user.clubs.where{ club_id == my{club.id} }.any?
    end

    can :new, Club
  end
end
