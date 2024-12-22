import { ColumnDef } from '@tanstack/react-table';
import { Checkbox } from '@/components/ui/checkbox';
import TableActionUser from './table-action';
import { User } from '@/types/admin/user';

export const userColumns: ColumnDef<User>[] = [
  {
    id: 'select',
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && 'indeterminate')
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: 'firstname',
    header: 'First Name',
    cell: ({ row }) => <div>{row.getValue('firstname')}</div>,
  },
  {
    accessorKey: 'lastname',
    header: 'Last Name',
    cell: ({ row }) => <div>{row.getValue('lastname')}</div>,
  },
  {
    accessorKey: 'email',
    header: 'Email',
    cell: ({ row }) => (
      <div className="font-mono truncate" title={row.getValue('email')}>
        {row.getValue('email')}
      </div>
    ),
  },
  {
    accessorKey: 'created_at',
    header: 'Created At',
    cell: ({ row }) => {
      const createdAt = row.getValue('created_at') as string;
      return <div>{new Date(createdAt).toLocaleString()}</div>;
    },
  },
  {
    accessorKey: 'updated_at',
    header: 'Updated At',
    cell: ({ row }) => {
      const updatedAt = row.getValue('updated_at') as string;
      return <div>{new Date(updatedAt).toLocaleString()}</div>;
    },
  },
  {
    accessorKey: 'deleted_at',
    header: 'Deleted At',
    cell: ({ row }) => {
      const deletedAt = row.getValue('deleted_at') as string | null;
      return (
        <div>{deletedAt ? new Date(deletedAt).toLocaleString() : '-'}</div>
      );
    },
  },
  {
    accessorKey: 'actions',
    header: () => <div className="text-right">Actions</div>,
    enableSorting: false,
    enableHiding: false,
    cell: ({ row }) => <TableActionUser user={row.original} />,
  },
];
