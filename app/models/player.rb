class Player < ActiveRecord::Base
  belongs_to :team

  has_many :scores

  has_many :player_games
  has_many :games, through: :player_games

  validates :name, presence: true
  validates :number, presence: true, uniqueness: { scope: :team }

  # Returns a special string that includes the name and number of the player
  #
  # @return [String] A special name for the player
  def number_name
    "#{self.number} - #{self.name}"
  end
end
