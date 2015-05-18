class Ability
  include CanCan::Ability

  def initialize(user)

    ###################################################
    ################## CLUBS ##########################
    ###################################################
    alias_action :new_admin, :create_admin, to: :update

    can :show, Club do |club|
      user.clubs.where{ id == my{club.id} }.any? || user.administrated_clubs.where{ id == my{club.id} }.any?
    end

    can :update, Club do |club|
      user.administrated_clubs.where{ id == my{club.id} }.any?
    end

    can :new, Club
    can :create, Club

    ###################################################
    ################## TEAMS ##########################
    ###################################################
    can :show, Team do |team|
      user.players.where{ team_id == my{team.id} }.any? ||
      user.administrated_clubs.joins{ teams }.where{ teams.id == my{team.id} }.any?
    end
  end
end
