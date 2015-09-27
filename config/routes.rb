UwrTournament::Application.routes.draw do
  root 'public#home'

  devise_for :users, controllers: { omniauth_callbacks: "users/omniauth_callbacks" }

  get '/statistics', to: 'public#statistics'

  resources :club_admins, only: [:new, :create, :destroy] do
    get 'confirm_destroy', on: :member
  end

  resources :club_join_requests, only: [:new, :create] do
    get 'accept', on: :member
    get 'decline', on: :member
  end

  resources :clubs, only: [:show, :new, :create, :edit, :update]

  resources :games, only: [:create, :new, :show, :index]

  resources :invitations, only: [:new, :create] do
    post 'accept', on: :member
    post 'decline', on: :member
  end

  resources :players, only: [:new, :create, :destroy] do
    get 'confirm_destroy', on: :member
  end

  resources :scores, only: [:create]

  resources :teams, only: [:new, :create, :show, :destroy, :edit, :update] do
    get 'confirm_destroy', on: :member
    get 'players', on: :member
  end

  resources :tournament_admins, only: [:new, :create, :destroy] do
    get 'confirm_destroy', on: :member
  end

  resources :tournament_invitations, only: [:new, :create, :destroy]

  resources :tournament_teams, only: [:new, :create]

  resources :tournaments, only: [:new, :create, :show]

  resources :user_clubs, only: [:destroy] do
    get 'confirm_destroy', on: :member
  end

  resources :users, only: [:show]
end
