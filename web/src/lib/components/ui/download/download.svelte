<script>
	import GoogleDownload from '$lib/components/design/button/google-download.svelte';
	import MegaDownload from '$lib/components/design/button/mega-download.svelte';

	const mega_download_link =
		'https://mega.nz/file/ut4jgSQb#zRM3Fk8RsSnoSgTu7sUfV632Bbt2v6UUt32Vyc3PWA4';
	const google_download_link =
		'https://drive.google.com/drive/folders/1fJqt9Z6hCyThYD4O65rfnZLnMEaLDrcy?usp=drive_link';
</script>

<div>
	<h1 class="text-xl font-semibold mb-5">Download the client from any one of the links</h1>
	<div class="flex gap-4 mb-5">
		<div class="my-5">
			<MegaDownload href={mega_download_link} />
		</div>
		<div class="my-5">
			<GoogleDownload href={google_download_link} />
		</div>
	</div>
	<p>After downloading the client, extract the files, and make sure your run C9Dark.exe</p>
</div>
