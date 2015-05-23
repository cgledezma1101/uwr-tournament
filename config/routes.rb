UwrTournament::Application.routes.draw do
  root 'games#new'

  devise_for :users

  get '/statistics', to: 'public#statistics'

  resources :clubs, only: [:show, :new, :create, :edit, :update] do
    get 'new_admin', on: :member
    post 'create_admin', on: :member
  end

  resources :games, only: [:create, :new, :show, :index]

  resources :scores, only: [:create]

  resources :teams, only: [:new, :create, :show] do
    get 'players', on: :member
  end

  resource :invitations, only: [:new, :create] do
    get 'accept', on: :member
    get 'decline', on: :member
  end
end
