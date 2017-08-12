class Ability
	include CanCan::Ability

	def initialize(user)
		alias_action :confirm_destroy, to: :destroy

		###################################################
		################## CLUBS ##########################
		###################################################
		can :create, Club

		can :read, Club do |club|
			user.clubs.where{ id == my{club.id} }.any? ||
			user.administrated_clubs.where{ id == my{club.id} }.any? ||
			user.pending_clubs.where{ id == my{club.id} }.any?
		end

		can :update, Club do |club|
			user.administrated_clubs.where{ id == my{club.id} }.any?
		end

		###################################################
		################## CLUB ADMINS ####################
		###################################################
		can :create, ClubAdmin do |club_admin|
			can? :update, club_admin.club
		end

		can :destroy, ClubAdmin do |club_admin|
			can? :update, club_admin.club
		end

		can :new, ClubAdmin

		###################################################
		################## CLUB JOIN REQUESTS #############
		###################################################
		can :new, ClubJoinRequest
		can :create, ClubJoinRequest

		can :accept, ClubJoinRequest do |club_join_request|
			user.administrated_clubs.where{ id == my{club_join_request.club_id} }.any?
		end

		can :decline, ClubJoinRequest do |club_join_request|
			user.administrated_clubs.where{ id == my{club_join_request.club_id} }.any?
		end

		###################################################
		################## GAME EVENTS ####################
		###################################################
		can :create, GameEvent do |game_event|
			can? :update, game_event.game
		end

		can :destroy, GameEvent do |game_event|
			can? :update, game_event.game
		end

		###################################################
		################## GAMES ##########################
		###################################################
		can :new, Game

		can :add_chronometer, Game

		can :create, Game do |game|
			can? :update, game.stage
		end

		can :destroy, Game do |game|
			(can? :update, game.stage) && !game.has_ended?
		end
		
		can :end, Game do |game|
			(game.status == Game::STATUS_STARTED) && (can? :update, game)
		end

		can :finalize, Game do |game|
			(game.status == Game::STATUS_STARTED) && (can? :update, game)
		end

		can :read, Game do |game|
			can? :read, game.stage
		end

		can :remove_score, Game do |game|
			can? :update, game
		end

		can :start, Game do |game|
			(game.status == Game::STATUS_READY) && (can? :update, game)
		end

		can :update, Game do |game|
			can? :update, game.stage
		end

		###################################################
		################## INVITATIONS ####################
		###################################################
		can :accept, Invitation do |invitation|
			user == invitation.user
		end

		can :decline, Invitation do |invitation|
			user == invitation.user
		end

		###################################################
		################## PLAYERS ########################
		###################################################
		can :new, Player

		can :create, Player do |player|
			can? :update, player.team
		end

		can :destroy, Player do |player|
			can? :update, player.team
		end

		###################################################
		################### SCORES ########################
		###################################################
		can :create, Score do |score|
			can? :update, score.game
		end

		###################################################
		################### STAGES ########################
		###################################################
		can :new, Stage

		can :create, Stage do |stage|
			can? :update, stage.tournament
		end

		can :destroy, Stage do |stage|
			(can? :update, stage.tournament) && !stage.has_ended_games?
		end

		can :read, Stage do |stage|
			can? :read, stage.tournament
		end

		can :update, Stage do |stage|
			can? :update, stage.tournament
		end

		###################################################
		################## TEAMS ##########################
		###################################################
		alias_action :add_player, to: :update

		can :new, Team

		can :create, Team do |team|
			can? :update, team.club
		end

		can :read, Team do |team|
			user.players.where{ team_id == my{team.id} }.any? ||
			user.administrated_clubs.joins{ teams }.where{ teams.id == my{team.id} }.any?
		end

		can :update, Team do |team|
			can? :update, team.club
		end

		can :destroy, Team do |team|
			can? :update, team.club
		end

		###################################################
		################## TOURNAMENTS ####################
		###################################################
		can :create, Tournament

		can :update, Tournament do |tournament|
			tournament.admins.where{ id == my{user.id} }.any?
		end

		can :read, Tournament do |tournament|
			user.administrated_tournaments.where{ id == my{tournament.id} }.any? ||
			user.clubs.joins{ tournaments }.where{ tournaments.id == my{tournament.id} }.any? ||
			user.administrated_clubs.joins{ tournaments }.where{ tournaments.id == my{tournament.id} }.any?
		end

		###################################################
		################## USER CLUBS #####################
		###################################################
		can :destroy, UserClub do |user_club|
			can? :update, user_club.club
		end

		###################################################
		####################### USERS #####################
		###################################################
		can :read, User do |authorized_user|
			user == authorized_user
		end
	end
end
