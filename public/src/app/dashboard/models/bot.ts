export interface Bot {
  id: number;
  name: string;
  build_result: string;
  webhook_secret: string;
  git_url: string;
}

export interface BuildStatusBuilt {
  type: 'built';
  created: string;
}

export interface BuildStatusBuilding {
  type: 'built';
}

export interface BuildStatusNone {
  type: 'none';
}

export type BuildStatus =
  | BuildStatusBuilding
  | BuildStatusBuilt
  | BuildStatusNone;

export interface Volume {
  bot_id: number;
  name: string;
  size: number;
  path: string;
}

export interface Config {
  bot_id: number;
  name: string;
  path: string;
  file: string;
  value: string;
}

export interface Env {
  bot_id: number;
  name: string;
  value: string;
}
