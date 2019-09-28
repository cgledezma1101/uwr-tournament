class PlayerGame < ApplicationRecord
	BLUE_TEAM = 'b'
	WHITE_TEAM = 'w'

	TEAM_COLORS = [BLUE_TEAM, WHITE_TEAM]

	belongs_to :player
	belongs_to :game

	validates :player, presence: true
	validates :game, presence: true
	validates :team_color, presence: true
	validate :valid_color

	private

	# Validates that the color of the team the player's in is valid
	def valid_color
		if !TEAM_COLORS.include?(self.team_color)
			self.errors.add(:team_color, t('player_game.errors.team_color', player: player.user.name))
		end
	end
end
