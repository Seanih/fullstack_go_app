import { useForm } from '@mantine/form';
import { useState } from 'react';
import { Button, Group, Modal } from '@mantine/core';

function AddTodo() {
	const [open, setOpen] = useState(false);

	const form = useForm({
		initialValues: {
			title: '',
			body: '',
		},
	});
	return (
		<div>
			<Modal opened={open} onClose={() => setOpen(false)} title='Create Todo'>
				Text
			</Modal>
			<Group position='center'>
				<Button fullWidth mb={12} onClick={() => setOpen(true)}>
					Add Task
				</Button>
			</Group>
		</div>
	);
}
export default AddTodo;
