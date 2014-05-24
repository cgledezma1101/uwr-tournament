class Player < ActiveRecord::Base
  belongs_to :team

  has_many :scores, dependent: :destroy

  has_many :player_games, dependent: :destroy
  has_many :games, through: :player_games

  validates :name, presence: true
  validates :number, uniqueness: { scope: :team }, allow_blank: true

  # Returns a special string that includes the name and number of the player
  #
  # @return [String] A special name for the player
  def number_name
    "#{self.number} - #{self.name}"
  end
end
