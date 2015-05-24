class User < ActiveRecord::Base
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable

   has_many :players

   has_many :user_clubs, dependent: :destroy
   has_many :clubs, through: :user_clubs

   has_many :club_admins
   has_many :administrated_clubs, through: :club_admins, source: :club

   has_many :invitations
   has_many :pending_clubs, through: :invitations, source: :club

   validates :name, presence: true
   validates :email, presence: true

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
end
