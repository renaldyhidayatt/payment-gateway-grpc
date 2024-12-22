import { ColumnDef } from '@tanstack/react-table';
import { Checkbox } from '@/components/ui/checkbox';
import TableActionTransfer from './table-action';
import { Transfer } from '@/types/admin/transfer';

export const transferColumns: ColumnDef<Transfer>[] = [
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
    accessorKey: 'transfer_from',
    header: 'Transfer From',
    cell: ({ row }) => <div>{row.getValue('transfer_from')}</div>,
  },
  {
    accessorKey: 'transfer_to',
    header: 'Transfer To',
    cell: ({ row }) => <div>{row.getValue('transfer_to')}</div>,
  },
  {
    accessorKey: 'transfer_amount',
    header: 'Transfer Amount',
    cell: ({ row }) => {
      const amount = row.getValue('transfer_amount') as number;
      const formatted = new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
      }).format(amount);
      return <div className="text-right font-medium">{formatted}</div>;
    },
  },
  {
    accessorKey: 'transfer_time',
    header: 'Transfer Time',
    cell: ({ row }) => {
      const time = row.getValue('transfer_time') as string;
      return <div>{new Date(time).toLocaleString()}</div>;
    },
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
    cell: ({ row }) => <TableActionTransfer transfer={row.original} />,
  },
];
