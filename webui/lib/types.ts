import { TableList } from '#components';
import { UITypes } from '~/utils/columnUtil';

export interface ITableSchema {
  name: string;
  type: string;
}

export interface ITableColumn {
  name: string;
  type: UITypes;
  options: any;
}

export type TableRowType = { id?: string | number } & { [key: string]: any }

export interface ITable {
  id?: number;
  name: string;
}

export interface ITab {
  label: string;
  key: string;
  favico: () => any;
  component: any;
  props: { [key: string]: any };
}

export interface IEndpoint {
  id: number;
  name: string;
  method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';
}

export type HTMLElementEvent<T extends HTMLElement> = Event & {
  target: T;
};

