class Ability
  include CanCan::Ability

  def initialize(user)

    ###################################################
    ################## CLUBS ##########################
    ###################################################
    can :show, Club do |club|
      user.clubs.where{ id == my{club.id} }.any? || user.administrated_clubs.where{ id == my{club.id} }.any?
    end

    can :new, Club
    can :create, Club
  end
end
