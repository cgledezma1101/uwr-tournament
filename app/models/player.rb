class Player < ActiveRecord::Base
  belongs_to :team

  has_many :scores

  has_many :player_games
  has_many :games, through: :player_games

  validates :name, presence: true
  validates :number, presence: true, uniqueness: { scope: :team }
end
