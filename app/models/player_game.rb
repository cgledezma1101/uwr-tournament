class PlayerGame < ActiveRecord::Base
  BLUE = 0
  WHITE = 1

  belongs_to :player
  belongs_to :game

  validates :player, presence: true, uniqueness: { scope: :game }
  validates :game, presence: true
  validates :color, presence: true
end
