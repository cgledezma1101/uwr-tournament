class Score < ApplicationRecord
	belongs_to :player
	belongs_to :game

	validates :player, presence: true
	validates :game, presence: true
end
