export interface OrbPagination<T> {
  limit: number;
  offset: number;
  order: string;
  total?: number;
  name?: string;
  tags?: string;
  dir?: 'desc' | 'asc';
  data: T[];
}

export interface NgxDatabalePageInfo {
  offset?: number;
  pageSize?: number;
  limit?: number;
  count?: number;
  name?: string;
  tags?: string;
}
