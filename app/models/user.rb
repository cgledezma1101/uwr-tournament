class User < ActiveRecord::Base
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

   # An ordered collection of the teams this user belongs to, which includes administrated and membership clubs
   #
   # @return [Array<Club>] All the clubs this user belongs to
   def all_clubs
     (self.clubs.includes(:teams).to_a + self.administrated_clubs.includes(:teams).to_a).uniq.sort do |club0, club|
       club0.name <=> club1.name
     end
   end

   # An ordered collection of the teams this user currently belongs to
   #
   # @return [Array<Team>] All the teams this user belongs to
   def all_teams
     Team.joins{ players }.where{ players.user_id == my{self.id} }.order(:name)
   end

   # Determines the amount of games this user has been involved in
   #
   # @return [Integer] Games the user has played
   def played_games
     Game.joins{ blue_players }.where{ blue_players.user_id == my{self.id} }.count +
     Game.joins{ white_players }.where{ white_players.user_id == my{self.id} }.count
   end

   # The amount of games where the team the user was playing in won
   #
   # @return [Integer] Amount of games the user has won
   def won_games
     Game.joins{ blue_players }.where{ (blue_players.user_id == my{self.id}) & (winning_color == 'blue') }.count +
     Game.joins{ white_players }.where{ (white_players.user_id == my{self.id}) & (winning_color == 'white') }.count
   end

   # The amount of games where the user participated and there was no victor
   #
   # @return [Integer] Amount of tied games where the user played
   def tied_games
     Game.joins{ blue_players }.where{ (blue_players.user_id == my{self.id}) & (winning_color == 'none') }.count +
     Game.joins{ white_players }.where{ (white_players.user_id == my{self.id}) & (winning_color == 'none') }.count
   end

   # The amount of games where the team the user was playing on lost
   #
   # @return [Integer] Amount of games the user has lost
   def lost_games
     Game.joins{ blue_players }.where{ (blue_players.user_id == my{self.id}) & (winning_color == 'white') }.count +
     Game.joins{ white_players }.where{ (white_players.user_id == my{self.id}) & (winning_color == 'blue') }.count
   end

   # Total goals scored by the user.
   #
   # @return [Integer] Amount of goals the user has scored
   def total_goals
     Score.joins{ player }.where{ player.user_id == my{self.id} }.count
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
