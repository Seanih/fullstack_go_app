import Head from 'next/head';
import Image from 'next/image';
import { Inter } from 'next/font/google';
import styles from '@/styles/Home.module.css';
import { Box } from '@mantine/core';
import useSWR from 'swr';
import { json } from 'stream/consumers';
import AddTodo from '@/components/AddTodo';

const inter = Inter({ subsets: ['latin'] });

export default function Home() {
	const endpoint = 'http://localhost:4000';

	const fetcher = async (url: string) => {
		let response = await fetch(`${endpoint}/${url}`);
		let data = await response.json();
		return data;
	};

	const { data, mutate } = useSWR('api/todos', fetcher);

	return (
		<>
			<Head>
				<title>Create Next App</title>
				<meta name='description' content='Generated by create next app' />
				<meta name='viewport' content='width=device-width, initial-scale=1' />
				<link rel='icon' href='/favicon.ico' />
			</Head>

			<AddTodo />
			<Box>{JSON.stringify(data)}</Box>
		</>
	);
}
