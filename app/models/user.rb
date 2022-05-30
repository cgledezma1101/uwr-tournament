class User < ApplicationRecord
    # Include default devise modules. Others available are:
    # :confirmable, :lockable, :timeoutable and :omniauthable
    devise :database_authenticatable,
                 :registerable,
                 :recoverable,
                 :rememberable,
                 :trackable,
                 :validatable
     devise :omniauthable, omniauth_providers: [:facebook, :google_oauth2]

     has_many :players

     has_many :user_clubs, dependent: :destroy
     has_many :clubs, through: :user_clubs

     has_many :club_admins, dependent: :destroy
     has_many :administrated_clubs, through: :club_admins, source: :club

     has_many :invitations, dependent: :destroy
     has_many :pending_clubs, through: :invitations, source: :club

     has_many :club_join_requests, dependent: :destroy
     has_many :pending_join_requests, through: :club_join_requests, source: :club

     has_many :tournament_admins, dependent: :destroy
     has_many :administrated_tournaments, through: :tournament_admins, source: :tournament

     validates :name, presence: true
     validates :email, presence: true, uniqueness: true

     # Tournaments in which the user will participate, or that the user administrates
     #
     # @return [Array<Tournament>] List of all the tournaments the user administrates or will participate in
     def active_tournaments
        admin_tournaments = Tournament.joins(:admins).where(tournament_admins: { user_id: self.id }).to_a
        player_tournaments = Tournament.joins(teams: :players).where(teams: { players: { user_id: self.id } }).to_a

        (admin_tournaments + player_tournaments).uniq
     end

     # An ordered collection of the teams this user belongs to, which includes administrated and membership clubs
     #
     # @return [Array<Club>] All the clubs this user belongs to
     def all_clubs
         (self.clubs.includes(:teams).to_a + self.administrated_clubs.includes(:teams).to_a).uniq.sort do |club0, club1|
             club0.name <=> club1.name
         end
     end

     # An ordered collection of the teams this user currently belongs to
     #
     # @return [Array<Team>] All the teams this user belongs to
     def all_teams
         Team.joins(:players).where( players: { user_id: self.id }).order(:name)
     end

     # The amount of games where the team the user was playing on lost
     #
     # @return [Integer] Amount of games the user has lost
     def lost_games
         Game
            .joins(player_games: :player)
            .where(player_games: { team_color: PlayerGame:: BLUE_TEAM }, winning_color: PlayerGame::WHITE_TEAM).or(
                Game
                    .joins(player_games: :player)
                    .where(player_games: { team_color: PlayerGame::WHITE_TEAM }, winning_color: PlayerGame::BLUE_TEAM)
            )
            .where(player_games: { players: { user_id: self.id } })
            .count
     end

     # Determines the amount of games this user has been involved in
     #
     # @return [Integer] Games the user has played
     def played_games
         Game.joins(:players).where(status: Game::STATUS_ENDED, players: { user_id: self.id }).count
     end

     # The amount of games where the user participated and there was no victor
     #
     # @return [Integer] Amount of tied games where the user played
     def tied_games
         Game.joins(:players).where( players: { user_id: self.id }, winning_color: Game::TIED_GAME).count
     end

     # Total goals scored by the user.
     #
     # @return [Integer] Amount of goals the user has scored
     def total_goals
         Score.joins(:player).where(players: { user_id: self.id }).count
     end

     # The amount of games where the team the user was playing in won
     #
     # @return [Integer] Amount of games the user has won
     def won_games
         Game
            .joins(player_games: :player)
            .where(player_games: { team_color: PlayerGame::BLUE_TEAM }, winning_color: PlayerGame::BLUE_TEAM).or(
                Game
                    .joins(player_games: :player)
                    .where(player_games: { team_color: PlayerGame::WHITE_TEAM }, winning_color: PlayerGame::WHITE_TEAM)
            )
            .where(player_games: { players: { user_id: self.id } })
            .count
     end

     # Allows a user to be logged in using an OmniAuth provider
     #
     # @param [Hash] auth Authentication information returned by the OmniAuth provider
     def self.from_omniauth(auth)
         where(email: auth.info.email).first_or_create do |user|
             user.email = auth.info.email
             user.name = auth.info.name
             user.password = Devise.friendly_token[0,20]
         end
     end
end
