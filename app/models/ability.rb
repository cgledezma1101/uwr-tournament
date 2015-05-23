class Ability
  include CanCan::Ability

  def initialize(user)
    alias_action :confirm_destroy, to: :destroy

    ###################################################
    ################## CLUBS ##########################
    ###################################################
    can :show, Club do |club|
      user.clubs.where{ id == my{club.id} }.any? || user.administrated_clubs.where{ id == my{club.id} }.any?
    end

    can :update, Club do |club|
      user.administrated_clubs.where{ id == my{club.id} }.any?
    end

    can :new, Club
    can :create, Club

    ###################################################
    ################## CLUB ADMINS ####################
    ###################################################
    can :new, ClubAdmin

    can :create, ClubAdmin do |club_admin|
      can? :update, club_admin.club
    end

    can :destroy, ClubAdmin do |club_admin|
      can? :update, club_admin.club
    end

    ###################################################
    ################## TEAMS ##########################
    ###################################################
    can :show, Team do |team|
      user.players.where{ team_id == my{team.id} }.any? ||
      user.administrated_clubs.joins{ teams }.where{ teams.id == my{team.id} }.any?
    end

    can :destroy, Team do |team|
      can? :update, team.club
    end

    ###################################################
    ################## UER CLUBS ######################
    ###################################################
    can :destroy, UserClub do |user_club|
      can? :update, user_club.club
    end
  end
end
