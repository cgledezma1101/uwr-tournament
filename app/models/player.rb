class Player < ActiveRecord::Base
  belongs_to :team
  belongs_to :user

  has_many :scores, dependent: :destroy

  has_many :player_games, dependent: :destroy
  has_many :games, through: :player_games

  validates :user, presence: true, uniqueness: { scope: :team }
  validates :number, uniqueness: { scope: :team }, allow_blank: true

  # Returns a special string that includes the name and number of the player
  #
  # @return [String] A special name for the player
  def number_name
    "#{self.number} - #{self.user.name}"
  end

  # Name of the user associated with this player
  #
  # @return [String] The name of the player
  def name
    self.number_name
  end

  # Name of the team the player belongs to
  #
  # @return [String] The name of the associated team
  def team_name
    self.team.name
  end

  # Email of the user associated with this player
  def email
    self.user.email
  end
end
