class Stage < ActiveRecord::Base
  belongs_to :tournament

  validates :tournament, presence: true
  validates :name, presence: true, uniqueness: { scope: :tournament }
  validate :name_not_blank

  private

  # Validates that the name that has been assigned to this stage is not null or empty
  def name_not_blank
    if self.name.blank?
      self.errors.add(:name, I18n.t('stage.errors.blank_name'))
    end
  end
end
