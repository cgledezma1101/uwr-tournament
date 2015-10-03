class Stage < ActiveRecord::Base
  belongs_to :tournament
  has_many :games

  validates :tournament, presence: true
  validates :name, presence: true, uniqueness: { scope: :tournament }
  validate :name_not_blank

  # Determines whether any of the games contained in this stage has ended
  #
  # @return [Boolean] Whether there's a finished game in this stage
  def has_ended_games?
    self.games.to_a.count{ |game| game.has_ended? } != 0
  end

  private

  # Validates that the name that has been assigned to this stage is not null or empty
  def name_not_blank
    if self.name.blank?
      self.errors.add(:name, I18n.t('stage.errors.blank_name'))
    end
  end
end
