TASK MANAGER API

#Giriş
Bu proje, basit bir görev yöneticisi uygulaması oluşturulmasına yardımcı olacak bir API sunar.


#Proje Açıklaması
Bu uygulama, kullanıcıların görevleri eklemelerine, listelemelerine, tamamlamalarına ve silmelerine olanak tanır. Veriler, hafif bir dosya tabanlı veri yapısında saklanabilir.

KURULUM

##kurulum

1. Projeyi klonlayın:

    ```bash
    git clone https://github.com/Fbulkaya/task_manager.git
    cd task_manager
    ```

2. Bağımlılıkları yükleyin:

    ```bash
    go mod tidy
    ```

3. Sunucuyu başlatın:

    ```bash
    go run main.go
    ```

4. Sunucu çalıştıktan sonra, `http://localhost:8080` adresinden API'ye erişebilirsiniz.

## Kullanım

### Görevleri Listeleme

- **GET** `/tasks`
- **Açıklama**: Tüm görevleri listeleyin.

### Görev Oluşturma

- **POST** `/tasks`
- **Açıklama**: Yeni bir görev oluşturun.
- **Örnek Gövde**:

    ```json
    {
        "name": "Yeni Görev",
        "completed": false
    }
    ```

### Görev Tamamla

- **PUT** `/tasks/{id}`
- **Açıklama**: Belirli bir görev tamamlandı olarak işaretleyin.

### Görev Silme

- **DELETE** `/tasks/{id}`
- **Açıklama**: Belirli bir görevi silin.

## Katkıda Bulunma

1. Projeyi fork edin.
2. Yeni bir dal oluşturun (`git checkout -b feature-branch`).
3. Değişikliklerinizi yapın ve test edin.
4. Dalınızı push edin (`git push origin feature-branch`).
5. Bir Pull Request oluşturun.

Lütfen katkıda bulunurken proje kurallarına ve kodlama standartlarına uyduğunuzdan emin olun.

